
# Declaring a class
class lotteryPlayers():
    def __init__(self):    # init method which has sone variables which is
    ## available to all the objects
        self.name = "Test"
        self.numbers = (10,23,1,234,2)

    def total(self):
        return sum (self.numbers)

# Creating a player object
player = lotteryPlayers()
print (player.name)
print (player.total())


### Different class in which an argument other than default "self" is used

class Student():
    def __init__(self, name, school): #init method which requires arguments
        self.name = name
        self.school = school
        self.mark = [25,]

    def print_list(self):
        print ("Name: {}".format(self.name))
        print ("List: {}".format(self.school))
        print ("Mark: {}".format(self.mark))

    def average(self):
        return (sum(self.mark)/len(self.mark))


anna = Student("Anna", "LMS") # passing argument and changing the value
anna.mark.append(23)  # appending the default empty list
anna.print_list()
print(anna.average())


####
