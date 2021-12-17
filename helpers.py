import subprocess
import sys

def printer_doc(path):
    if sys.platform == 'win32':
        args = '"C:\\\\Program Files\\\\gs\\\\gs9.23\\\\bin\\\\gswin64c" ' \
            '-sDEVICE=mswinpr2 ' \
            '-dBATCH ' \
            '-dNOPAUSE ' \
            '-dFitPage ' \
            '-sOutputFile="%printer%myPrinterName" '
        ghostscript = args + str(path).replace('\\', '\\\\')
        subprocess.call(ghostscript, shell=True)
    else:
        print(f"Print documento {path}")