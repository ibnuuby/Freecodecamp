# Python For-loop

# for letter in 'Hello':
# 	print(letter)

# Python for-loops and lists

# Loop over the list
# mylist = [1, 'a', 'Hello']
# for item in mylist:
# 	print(item)

# Loop single item
# mylist = [1, 'a', 'Hello']
# for item in mylist:
# 	print(mylist[1])

# Python While-Loop
# i = 1
# while i <= 5:
# 	print(i)
# 	i = i+1


# Infinite Loop
# while True:
# 	print("True")

# Python Function: The Basics Of Code Reuse
# print("Hi, am learn programming python")

# Creating a Python function
# def say_hi():
# 	print("hi")
# say_hi()

# Python function parameters and arguments
# def say_hi(name):
#     print("Hi", name)
# say_hi('Inu')

# Python function with multiple arguments
# def welcome (name, location):
#     print("Hi", name, "welcome to", location)
# welcome("erik", "this is tutorial")

# Returning from a function
# def say_hi(name):
# 	print('Hi, ' + name)
# print("Lets greet the entire world")
# say_hi('world')


# Returning values
# def add(a, b):
# 	return a + b
# result = add(4,8)
# print(result)

# Empty return statement
# def optional_greeter(name):
#     if name.startswith('X'):
#         # We don't greet people with weird names :p
#         return
    
#     print('Hi there, ', name)
# optional_greeter('Xander')

# def optional_greeter(name):
#     if name.startswith('X'):
#         # We don't greet people with weird names :p
#         pass
#     else:
#         print('Hi there, ', name)
# optional_greeter('Xander')

# Variable scope
# def say_hi():
#     print('Hi', name)
#     answer = "Hi"
# name = 'John'
# say_hi()

# Default values and named parameters
# >>> def welcome(name='learner', location='this tutorial'):
# ...     print("Hi", name, "welcome to", location)
# ...
# >>> welcome()
# Hi learner welcome to this tutorial
# >>> welcome(name='John')
# Hi John welcome to this tutorial
# >>> welcome(location='this epic tutorial')
# Hi learner welcome to this epic tutorial
# >>> welcome(name='John', location='this epic tutorial')
# Hi John welcome to this epic tutorial

#Resource
https://python.land/introduction-to-python/functions
