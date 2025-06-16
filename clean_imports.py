import argparse
import os
import re

def delete_import(file, package):
    with open(file, 'r') as f:
        content = f.read()

    patron = re.compile(
        r'import\s*\(\s*([^)]*?)\s*\)|import\s*"([^"]+)"',
        re.DOTALL
    )

    def update_import(match):
        if match.group(1):
            lines = [line.strip() for line in match.group(1).split('\n') if line.strip()]
            filter = [line for line in lines if f'"{package}"' not in line]
            return f'import (\n    ' + '\n    '.join(filter) + '\n)'
        elif match.group(2):
            if f'"{package}"' in match.group(2):
                return match.group(0)
            else:
                return ''
        return match.group(0)

    with open(file, 'w') as f:
        f.write(re.sub(r'\n\n+', '\n\n', patron.sub(update_import, content).strip())
)

parser = argparse.ArgumentParser()
parser.add_argument('pkg')
args = parser.parse_args()

for root, e, files in os.walk("."):
    for file in files:
        if file.endswith(".go"):
            path = os.path.join(root, file)
            delete_import(path, args.pkg[4:])
