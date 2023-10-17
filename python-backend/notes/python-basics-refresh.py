# Basic data types
name: str = "John" # strings
number: int = 10 # integer
decimal: float = 3.1416 # floating point number
isTrue: bool = True # boolean 

# data structures
digits = [1,2,3,4,5,6,7,8,9] # list, indexed values
vowels = {"a", "e", "i", "o", "u"} # set, unique values
result = ("A", "B", "C") # tuples (useful for multiple returns) indexed inmutable values
option = {"a": "piedra", "b": "papel", "c": "tijera"} # dictionaries, map keys to values

# typed data structures
l: list[int] = [1,2,3]
s: set[str] = {"a", "e", "i", "o", "u"}
t: tuple[int, int, int] = (1,2,3)
d: dict[str, str] = {"name": "Akon", "age": "50"}

# multiple types
m: list[int | float] = [1,2.5,3]

# match
name = "Carl"
match name[0]:
    case "A":
        print("Name starts with A")
    case "B":
        print("Name starts with B")
    case _:
        print("Name dont't starts with A or B")

# functions and default values
def greet(name="world"):
    return f"hello {name}!"

# functions with dynamic arguments
def dyn_args(*args, **kwargs):
    # args = standard arguments (tuple) useful to pass multiple arguments
    # kwargs = key arguments (dictionary) useful to pass specific arguments
    print("args: ", args)
    print("kwargs: ", kwargs)

# list comprehension
odd_numbers = [n for n in range(0,11) if n % 2 == 0 ]
for n in odd_numbers:
    if n != 0:
        print(n)

# generators = can be used only one time
generator = (n for n in range(0, 11) if n % 2 == 0)
even = list(generator)
even_bis = list(generator)
print(even)             # [0, 2, 4, 6, 8, 10]
print(even_bis)         # []

# OOP
class Animal:
    def hablar(self):
        print("Animal hace ruido")

class Perro(Animal):
    def hablar(self):
        print("Perro ladra")

class Gato(Animal):
    pass  # No se define el método hablar en la subclase Gato

animal = Animal()
perro = Perro()
gato = Gato()

animal.hablar()  # Animal hace ruido
perro.hablar()   # Perro ladra
gato.hablar()    # Animal hace ruido (hereda la lógica de la superclase)

# Sobreescribir metodos y uso de super() para conservar lógica
class Animal2:
    def __init__(self, nombre):
        self.nombre = nombre

class Perro2(Animal2):
    def __init__(self, nombre, raza):
        super().__init__(nombre)  # Llama al constructor de la superclase
        self.raza = raza

perro = Perro2("Fido", "Labrador")
print("Nombre:", perro.nombre)  # Nombre: Fido
print("Raza:", perro.raza)      # Raza: Labrador


# Estructura basica
def main():
    msg = greet("Silas")
    print(msg)
    dyn_args(1,2,3,a=1,b=2)

if __name__ == "__main__":
    main()
