import pytest
from app import app

@pytest.fixture
def client():
    with app.test_client() as client:
        yield client

def test_duplicates_no_text(client):
    res = client.get('/duplicate')
    assert res.status_code == 400
    assert res.get_json() == {"error": True, "string": "Missing 'text' parameter", "answer": ""}

def test_duplicates_empty_text(client):
    res = client.get('/duplicate?text=')
    assert res.status_code == 400
    assert res.get_json() == {"error": True, "string": "Missing 'text' parameter", "answer": ""}

def test_duplicates_space_text(client):
    res = client.get('/duplicate?text=    ')
    assert res.status_code == 200
    assert res.get_json() == {"error": False, "string": "Found duplicates", "answer": ""}

def test_duplicates_true(client):
    res = client.get('/duplicate?text=hello+world+hello')
    assert res.status_code == 200
    assert res.get_json() == {"error": False, "string": "Found duplicates", "answer": "hello"}

def test_duplicates_false(client):
    res = client.get('/duplicate?text=hello+world+cool')
    assert res.status_code == 200
    assert res.get_json() == {"error": False, "string": "Found duplicates", "answer": ""}

def test_duplicates_special_characters(client):
    res = client.get('/duplicate?text=hello%2C+world!+hello.')
    assert res.status_code == 200
    assert res.get_json() == {"error": False, "string": "Found duplicates", "answer": "hello"}

def test_duplicates_breaking_word_special(client):
    res = client.get('/duplicate?text=hello%2C+world!+hell*o.')
    assert res.status_code == 200
    assert res.get_json() == {"error": False, "string": "Found duplicates", "answer": ""}
