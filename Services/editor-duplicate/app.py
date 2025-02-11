from flask import Flask, request, jsonify
import re

app = Flask(__name__)

def find_duplicates(text):
    words = re.split(r'\W+', text.lower())
    seen = set()
    duplicates = set()
    for word in words:
        if word in seen:
            duplicates.add(word)
        else:
            seen.add(word)
    return ' '.join(duplicates)

@app.route('/duplicate', methods=['GET'])
def process():
    text = request.args.get('text', '')
    if not text:
        return jsonify({"error": True, "string": "Missing 'text' parameter", "answer": ""}), 400
    duplicates = find_duplicates(text)
    return jsonify({"error": False, "string": "Found duplicates", "answer": duplicates})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8030)

