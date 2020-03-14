
building_list = [x * 3 for x in range(5) ] ### Multiplies every number from 1 to 5 by 3 and creates a list
print (building_list)

building_list_condtional = [x for x in range(20)  if x%2 == 0] ### Creates a list of numbers divisible by 2 
print(building_list_condtional)


#Use list comprehension

def who_do_you_know():
    input_list = input ("give me a list of people sperated by comma")
    list_names = input_list.split(",")
    list_names_nospace = [i.strip().lower() for i in list_names ] # Remove white spaces
    return list_names_nospace

def ask_user():
    name = input ("give me a name")
    if name in who_do_you_know():
        print("I know {}".format(name))

ask_user()
