import sys
from parse_input import parseInput

def main() -> None:
	try:
		jsonData = parseInput(sys.stdin)
		print(jsonData)
	except Exception as error:
		print("Error:", error)
		sys.exit(-1)

if __name__ == "__main__":
	main()
