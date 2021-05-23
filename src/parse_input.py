import json
import datetime
from typing import TextIO

def parseInput(stdin: TextIO) -> dict:
	line = stdin.readline()
	jsonData = json.loads(line)
	now = datetime.datetime.now(datetime.timezone.utc)
	jsonData['date'] = now.isoformat()
	return jsonData
