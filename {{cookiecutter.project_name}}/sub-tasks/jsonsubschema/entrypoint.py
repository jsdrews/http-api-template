import sys

import yaml
import json
from jsonsubschema import isSubschema


usage = """
jsonsubschema checks if one JSON schema is a subschema (subtype) of another.

For any two JSON schemas s1 and s2, s1 <: s2 (reads s1 is subschema/subtype of s2) if every JSON document instance that validates against s1 also validates against s2.
"""


def load(filepath: str) -> dict:
	if filepath.endswith(".yaml"):
		with open(filepath) as file:
			return yaml.safe_load(file)
	elif filepath.endswith(".json"):
		with open(filepath) as file:
			return json.load(file)
	else:
		raise ValueError("Only supports .yaml and .json files")


if __name__ == "__main__":
	if len(sys.argv) < 3:
		raise Exception(
			f"usage: entrypoint.py jsonschema1 jsonschema2\n{usage}"
		)
	s1 = load(sys.argv[1])
	s2 = load(sys.argv[2])
	backwards_compatible = isSubschema(s1, s2)
	print(f's1 is subset of s2: {backwards_compatible}')
	if not backwards_compatible:
		sys.exit(1)
