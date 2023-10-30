
var first = "Hello";
var second = "World!";

Console.WriteLine("======= Composite Formating ======");
Console.WriteLine("{0} {1}", first, second);
Console.WriteLine("{1} {0}", first, second);
Console.WriteLine("{0} {0} {0}", first, second);

Console.WriteLine("\n======= String Interpolation ======");
Console.WriteLine($"{first} {second}");
Console.WriteLine($"{second} {first}");

Console.WriteLine("\n======= Formating Currency ======");
decimal price = 120.45m;
int discount = 50;
Console.WriteLine($"Price: {price:C}, Discount: {discount:C}");

Console.WriteLine("\n======= Formating Decimals ======");
decimal measurement = 120_456.48988m;
Console.WriteLine($"Measurement (2 spaces): {measurement:N2}");
Console.WriteLine($"Measurement (4 spaces): {measurement:N4}");


Console.WriteLine("\n======= Formating Percentages ======");
decimal tax = .1745m;
Console.WriteLine($"Tax percentage {tax:P2}");


Console.WriteLine("\n======= Formating Percentages ======");
decimal Length = 1.1234m;
decimal Percentage = .1234m;
int Money = 1234;
var msg = String.Format("Percentage: {1:P2}, Length: {0:N2}, Money: {2:C}", Length, Percentage, Money);
Console.WriteLine(msg);


