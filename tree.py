num = int(input("input number of rows: ")) 
stars = 1

while num > 0:
  num -= 1

  print(" "*num + "*"*stars)
  stars += 2