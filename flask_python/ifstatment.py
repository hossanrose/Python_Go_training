variable = True

if variable: #Dont have to write a test explicitly
    print ("Test")

know = ["test","Hossan","Varada"]
person = input("Type a name")
if person in know:
    print("I know {}".format(person))
else:
    print("I dont know {}".format(person))

def who_do_you_know():
    input_list = input ("give me a list of people sperated by comma")
    list_names = input_list.split(",") # Make a list by elements seperated by commas
    list_names_nospace =[]
    for i in list_names:
        list_names_nospace.append(i.strip()) # Remove white spaces
    return list_names_nospace

def ask_user(list_names):
    name = input("give me a name")
    if name in list_names:
        print("I know {}".format(name))

list_names=who_do_you_know()
ask_user(list_names)
