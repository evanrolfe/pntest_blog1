import os
from http.server import BaseHTTPRequestHandler, HTTPServer

import requests

host_name = os.environ['HOST_NAME']
port = os.environ['PORT']

RESPONSE_HTML = """
<html>
    <head><title>PnTest</title></head>
    <body>
        <p>This is an example web server.</p>
        <p>Request path: %s</p>
    </body>
</html>
"""

class ServerHandlers(BaseHTTPRequestHandler):
    def do_GET(self):
        response = requests.get("http://pntest.io" + self.path)
        print("Response from pntest.io: ", response.status_code)

        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes(RESPONSE_HTML % self.path, "utf-8"))

if __name__ == "__main__":
    server = HTTPServer((host_name, int(port)), ServerHandlers)
    print("Server starting at http://%s:%s" % (host_name, port))
    server.serve_forever()
