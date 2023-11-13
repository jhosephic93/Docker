from flask import Flask
import os

app = Flask(__name__)

# Configuración externa (por ejemplo, desde un archivo .env)
app.config['DEBUG'] = os.getenv('FLASK_DEBUG', 'True') == 'True'
app.config['PORT'] = os.getenv('FLASK_PORT', 8080)

@app.route("/my-container")
def awsecs():
    return "Hi There... I'm running on a machine"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=app.config['PORT'], debug=app.config['DEBUG'])
