import requests
import json
import os
import sys
from argparse import ArgumentParser


class Leetcode:
    url = 'https://leetcode.com'

    def __init__(self, session_id):
        self.session = requests.Session()
        self.session.cookies = requests.sessions.cookiejar_from_dict({
            'LEETCODE_SESSION': session_id
        })

    def raw(self, method, uri, params=None, data=None):
        headers = {
            'content-type': 'application/json',
        }
        return self.session.request(method, self.url + uri, headers=headers, params=params, data=data)

    def call_graphql(self, data):
        response = self.raw('POST', '/graphql', data=data)
        return response.json()


def select_useful_data(src, lang='Go'):
    result = {}

    question = src['data']['question']

    result['slug'] = question['titleSlug']
    result['meta_name'] = json.loads(question['metaData'])['name']
    result['meta_name_normal'] = result['meta_name'][0].upper() + result['meta_name'][1:]
    for snippet in question['codeSnippets']:
        if snippet['lang'] == lang:
            result['code'] = snippet['code']

    return result


def get_question(leetcode_cls, slug):
    query = '''
    query questionData($titleSlug: String!) {
        question(titleSlug: $titleSlug) {
            title titleSlug similarQuestions exampleTestcases
            codeSnippets {
                lang langSlug code __typename
            }
            sampleTestCase
            metaData
            __typename
        }
    }
    '''
    payload = {
        'operationName': 'questionData',
        'variables': {
            'titleSlug': slug,
        },
        'query': query
    }
    resp = leetcode_cls.call_graphql(
        json.dumps(payload)
    )
    return select_useful_data(resp)


def write_to_file(file_path, data):
    with open(file_path, 'w') as file:
        file.write(data)


def get_template(tpl_path, lang):
    tpl_file = os.path.join(tpl_path, lang.lower() + '.tpl')
    with open(tpl_file, 'r') as file:
        return file.read()


if __name__ == '__main__':
    parser = ArgumentParser()
    parser.add_argument('-l', required=True, dest='language', default='Go', help='Language')
    parser.add_argument('-e', dest='expansion', default='go', help='Code file expansion')
    parser.add_argument('-s', dest='suffix', default='_test', help='Code file suffix')
    parser.add_argument('-d', dest='dst', help='Solutions destination path')
    parser.add_argument('-t', dest='templates', help='Templates path')
    parser.add_argument('slug', help='Slug')
    args = parser.parse_args()

    template = get_template(args.templates, args.language)
    slug = args.slug
    if slug.startswith('http'):
        for i in reversed(slug.split('/')):
            if i:
                slug = i
                break
    lc_session_id = os.getenv('LEETCODE_SESSION_ID')
    if not lc_session_id:
        print('env LEETCODE_SESSION_ID is not set')
        sys.exit(1)
    lc = Leetcode(lc_session_id)
    quest_data = get_question(lc, slug)
    code = template.format(**quest_data)
    if args.dst:
        code_file = os.path.join(args.dst, quest_data['slug'] + f'{args.suffix}.{args.expansion}')
        write_to_file(code_file, code)
    else:
        print(code)
