my_variable = "hello"
my_list =[23, 12 ,2,45,6]

#For loop
for i in my_variable: #iteration can be done with sets, tuples, list and string
    print (i)

for n in my_list:
    print (n)

#While loop
user_wants_number = True

while user_wants_number == True:
    print ("Yo I'm good")
    user_input = input("Give me an answer y/n:")
    if user_input == "n":
        user_wants_number = False
