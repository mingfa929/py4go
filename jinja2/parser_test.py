import parser
# import pytest
# import jinja2.parser

# jinja2.parse("dddd")

def test_jinja2():
    print("------->test_jinja2")
    inputs=parser.parse("dddd")
    print(inputs)
    assert 1 == 1

if __name__ == '__main__':
    # pytest.main(["-v", "-s", "jinja2_test.py"])
    test_jinja2()