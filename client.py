import click
import requests
import config
from pathlib import Path

@click.group()
@click.pass_context
def cli(ctx):
    pass


@cli.command("print")
@click.argument('file')
@click.option('--printer', '-p', default="default")
def print_(file, printer):
    url = f"{config.URL_HOST}/print?printer={printer}"
    path = Path(file).absolute()
    files = {'file': path.open('rb')}
    r = requests.post(url, files=files)
    if r.status_code != 200:
        print(r.content)
    else:
        data = r.json()
        print(data)

if __name__ == '__main__':
    cli(obj={})