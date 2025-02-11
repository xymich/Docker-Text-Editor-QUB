import pytest
from app import find_duplicates 
# test cases - unit testing
def test_duplicates_no_duplicates():
    text = "Hello world"
    assert find_duplicates(text) == ""

def test_duplicates_with_duplicates():
    text = "Hello world hello"
    assert find_duplicates(text) == "hello"

def test_duplicates_ignore_case():
    text = "Hello world hello"
    assert find_duplicates(text) == "hello"

def test_duplicates_with_special_characters():
    text = "Hello, world! Hello."
    assert find_duplicates(text) == "hello"

def test_duplicates_empty_string():
    text = ""
    assert find_duplicates(text) == ""

def test_duplicates_spaced_string():
    text = "   "
    assert find_duplicates(text) == ""

def test_duplicates_breaking_word_special():
    text = "/duplicate?text=hello%2C+world!+hell*o."
    assert find_duplicates(text) == ""