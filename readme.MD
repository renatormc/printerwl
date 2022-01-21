# .env
```
URL_HOST=http://localhost:5000
PASSWORD=asdfg
DEFAULT_PRINTER=printername
PORT=8001
TLS_ENABLED=false
```

# Dependencies

## Windows

It's necessary to install PDFtoPrinter first:
<http://www.columbia.edu/~em36/pdftoprinter.html/>

# Create link

```
sudo ln -s "$(pwd)/exedir/rprinter" /usr/local/bin/rprinter
```