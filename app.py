from flask import Flask, request, jsonify
import config
from uuid import uuid4
import helpers

app = Flask(__name__)

@app.route("/print", methods=("POST",))
def index():
    printer = request.args.get("printer") 
    if not printer or printer == 'default':
        printer = config.DEFAULT_PRINTER
    file = request.files['file']
    path = config.tempfolder / f"{uuid4()}.pdf"
    file.save(str(path))
    helpers.printer_doc(path)
    print(printer)
    return jsonify({'path': str(path)})
