import os
from pathlib import Path
import tempfile
from dotenv import load_dotenv
load_dotenv()

app_dir = Path(os.path.dirname(os.path.realpath(__file__)))

URL_HOST = os.getenv("URL_HOST")
PASSWORD = os.getenv("PASSWORD")
DEFAULT_PRINTER = os.getenv("DEFAULT_PRINTER")

tempfolder = Path(tempfile.gettempdir())


