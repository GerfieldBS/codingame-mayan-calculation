# CodinGame - Mayan Calculation
https://www.codingame.com/training/medium/mayan-calculation

Upon discovering a new Maya site, hundreds of mathematics, physics and astronomy books have been uncovered.

The end of the world might arrive sooner than we thought, but we need you to be sure that it doesn't!

Thus, in order to computerize the mayan scientific calculations, you're asked to develop a program capable of performing basic arithmetical operations between two mayan numbers.
 	Rules
The mayan numerical system contains 20 numerals going from 0 to 19. Here's an ASCII example of their representation:

| line | 0 | 1     | 2     | 3     | 4     | 5     | 6     | 7     | 8     | 9     | 10    | 11    | 12    | 13    | 14    | 15    | 16    | 17    | 18    | 19    | 
|------|------| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| 1    | .oo. | o... | oo.. | ooo. | oooo | .... | o... | oo.. | ooo. | oooo | .... | o... | oo.. | ooo. | oooo | .... | o... | oo.. | ooo. | oooo | 
| 2    | o..o | .... | .... | .... | .... | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | 
| 3    | .oo. | .... | .... | .... | .... | .... | .... | .... | .... | .... | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | ____ | 
| 4    | .... | .... | .... | .... | .... | .... | .... | .... | .... | .... | .... | .... | .... | .... | .... | ____ | ____ | ____ | ____ | ____ | 

A mayan number is divided into vertical sections. Each section contains a numeral (from 0 to 19) representing a power of 20. The lowest section corresponds to 200, the one above to 201 and so on.

Thereby, the example below is : (12 x 20 x 20) + (0 x 20) + 5 = 4805

| Mayan | Digit position | Formula   |
| --- |----------------|-----------|
| oo.. | 1              | 12 x 20^2 |
| ____ | 1              |           |
| ____ | 1              |           |
| .... | 1              |           |
| .oo. | 2              | 0 x 20^1 |
| o..o | 2              |           |
| .oo. | 2              |           |
| .... | 2              |           |
| .... | 3              | 5 x 20^0 |
| ____ | 3              |           |
| .... | 3              |           |
| .... | 3              |           |

To spice the problem up, the mayans used several dialects, and the graphical representation of the numerals could vary from one dialect to another.
 
The representation of the mayan numerals will be given as the input of your program in the form of ASCII characters. You will have to display the result of the operation between the two given mayan numbers. The possible operations are *, /, +, -