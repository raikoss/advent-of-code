import os

class FinishedIntcode(Exception):
  pass

def get_opcode_result(opcode, value1, value2):
  if opcode == 1:
    return value1 + value2
  if opcode == 2:
    return value1 * value2
  if opcode == 99:
    return None

def run_opcode(intcode, position):
  opcode = intcode[position]

  if opcode == 99:
    raise FinishedIntcode(intcode)

  position += 1
  value1_position = intcode[position]
  position += 1
  value2_position = intcode[position]
  position += 1
  destination = intcode[position]
  position += 1

  value1 = intcode[value1_position]
  value2 = intcode[value2_position]

  value = get_opcode_result(opcode, value1, value2)
  intcode[destination] = value

def run_intcode(intcode):
  position = 0

  try:
    # Yeehaaw
    while True:
      run_opcode(intcode, position)
      position += 4
      
  except FinishedIntcode:
    print("New intcode", intcode)


absolute_path = os.path.dirname(__file__)

with open(absolute_path + "./input.txt", "r") as file:
  intcode_string = file.readline()
  print("Running intcode", intcode_string)

intcode = list(map(lambda int_string: int(int_string), intcode_string.split(',')))

# Assignment tells you to
intcode[1] = 12
intcode[2] = 2

run_intcode(intcode)