class Car:
    def __init__(self):
        self.brand = None
        self.engine = None
        self.tires = None
        self.color = None
        self.country = None

    def __str__(self):
        return f"Car: {self.color} {self.brand} with {self.engine} engine and {self.tires} tires"


def car_factory(car_type):
    if car_type == "lexus":
        lexus = Car()
        lexus.brand = "lexus"
        lexus.engine = "v6"
        lexus.tires = "4"
        lexus.color = "blue"
        lexus.country = "Japan"
        return lexus
    elif car_type == "mercedes":
        mercedes = Car()
        mercedes.brand = "benz"
        mercedes.engine = "v8"
        mercedes.tires = "6"
        mercedes.color = "black"
        mercedes.country = "germany"
        return mercedes
    return None


# Lexus ES - V6 - 4 - blue Japan
# Mercedes-AMG - G 63 - V-8  - 6 -black germany

mercedes = car_factory("mercedes")
lexus = car_factory("lexus")
print(mercedes.__str__())
print(lexus.__str__())
