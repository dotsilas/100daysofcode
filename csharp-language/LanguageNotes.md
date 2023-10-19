
# C#
- case sensitive
- semicoloms
- literal values == constant value unchangeable
- basic datatypes
```cs
// char == single alphanumerical value (single quotes)
char miLetter = 'A';

// strings == secuences of characters (double quotes)
string myString = "This is a string";

// int == whole numbers
int myIntegerNumber = 123;

// floating-points numbers
// float == 6-9 digits precision
Console.WriteLine(.25F); // float

// double == 15-17 digits precision
Console.WriteLine(2.625); // double, default 

// decimal == 28-29 digits precision
Console.WriteLine(.2512341234m); // float

// boolean == true or false values
Console.WriteLine(true);
Console.WriteLine(false);
```
- Variables
   - namming conventions;
      - variables: camelCase
   - declaration
      datatype + name;
   - declaration and setting value
      - datatype + name = value;
   - typeinference
      - var name = value
- String formating
   - \" to use "" inside of a string
   - \n new line
   - \t tab space
   - \\ backslash
   - \u use utf characters
   - verbatim strings == keep all spaces and characters without backslash notation (raw string)
      - @"my verbatim string"
   - string interpolation
      - $"{value1} and {value2}"
   - verbatim + string interpolation
      - $@"http://google.com/{query}"
- Operations
   - can use () to indicate that + inside ins't concat... 
   - for division, default is integer division, you need to use at least one float number in division for real division
   - % modulus operator
   - PEMDAS
   - incremental operators: number+=, number++, number = number + 1
   - increment before and after value
      - Console.WriteLine(++value); // prints value + 1
      - Console.WriteLine(value++); // prints value and adds + 1 to value
- For... // TO-DO: tengo que aprender diferencia entre for y foreach, cuando es mejor cada uno
   - foreach don't allow to change the value of the iterable variable
   - classical for, allows to change the value of the iterable variable
   - break breaks iteration
- Read from console
   - Console.ReadLine();
- Nullable values
   - int userNumber; // automatically assign '0' as default value
   - int? userNumber; // assign null as predetermined value, so is null until initialization
- 
