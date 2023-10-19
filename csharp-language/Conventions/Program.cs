/*
// char userOption;
// int gameScore;
// float particlesPerMillion;
// bool processedCustomer;

string firstName = "Bob";
// int widgetsPurchased = 7;
// Testing a change to the message.
int widgetsSold = 7;
Console.WriteLine($"{firstName} sold {widgetsSold} widgets.");
// Console.WriteLine($"{firstName} purchased {widgetsPurchased} widgets.");
*/

// dificult to read code
static void DificultCode()
{
    Random dice = new Random();
    int roll1 = dice.Next(1, 7);
    int roll2 = dice.Next(1, 7);
    int roll3 = dice.Next(1, 7);
    int total = roll1 + roll2 + roll3;
    Console.WriteLine($"Dice roll: {roll1} + {roll2} + {roll3} = {total}");
    if ((roll1 == roll2) || (roll2 == roll3) || (roll1 == roll3)) {
        if ((roll1 == roll2) && (roll2 == roll3)) {
            Console.WriteLine("You rolled triples!  +6 bonus to total!");
            total += 6; 
        } else {
            Console.WriteLine("You rolled doubles!  +2 bonus to total!");
            total += 2;
        }
    }
}

// more readable spacing
static void DificultCode()
{
    Random dice = new Random();

    int roll1 = dice.Next(1, 7);
    int roll2 = dice.Next(1, 7);
    int roll3 = dice.Next(1, 7);

    int total = roll1 + roll2 + roll3;
    Console.WriteLine($"Dice roll: {roll1} + {roll2} + {roll3} = {total}");

    if ((roll1 == roll2) || (roll2 == roll3) || (roll1 == roll3))
    {
        if ((roll1 == roll2) && (roll2 == roll3))
        {
            Console.WriteLine("You rolled triples!  +6 bonus to total!");
            total += 6; 
        }
        else
        {
            Console.WriteLine("You rolled doubles!  +2 bonus to total!");
            total += 2;
        }
    }
}
