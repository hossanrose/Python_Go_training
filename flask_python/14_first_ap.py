from flask import Flask

app = Flask(__name__) ## Creating an app

@app.route("/")
def test():
    return "Test"

app.run(port=1000)
