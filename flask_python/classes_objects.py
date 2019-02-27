
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
