import math
import os

def calculate_fuel_recurive(mass):
  total_fuel = 0
  fuel = calculate_fuel(mass)

  while fuel > 0:
    print(fuel)
    total_fuel += fuel
    fuel = calculate_fuel(fuel)
  
  return total_fuel

def calculate_fuel(mass):
  fuel = math.floor(int(mass) / 3) - 2
  return fuel

absolute_path = os.path.dirname(__file__)

with open(absolute_path + "./input.txt", "r") as file:
  total_fuel = 0

  for line in file:
    fuel = calculate_fuel_recurive(line)
    total_fuel += fuel

print(total_fuel)