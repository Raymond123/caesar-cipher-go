# caesar-cipher-go

1st assignment written in go.

part 1: 
Created caesar cipher function that take a string and shift input and gives a casaer cipher of the string by shifting each character the 'shift' amount in the alphabet. If the shift goes past 'Z' then it loops back around to 'A' and continues adding the rest of 'shift' from there.

part 2:
uses the same caesar cipher code but in a different function that takes a list of strings, 'shift' value, and 2 channels for syncing. this function is then run as a go routine and synced in the main function.

part 3:
uses the same list caesar function as part 2, but runs 3 separate go routines each taking 1/3rd of the message array as their msg input. uses the same syncing as part 2.

caesar function specifics:
in order to shift each character, first the string is seperated into an array of runes. then each rune that is an alphabetic character is converted to its integer value and has shift added on. the new integer is then converted back into a rune and added to a new rune array. if the character after adding shift is no longer an alphabetic character then the function subtracts 26 (# of letters in english) from the runes integer value.
when going to print, the rune array is converted to a single string and each letter is capitalized. 

parts 2/3 syncing:
parts 2 and 3 are synced using 2 channels, a string channel and an int channel. when a message is finished being encrypted it is sent to the string channel, 'item', and then the main function reads that channel and prints the result. at the same time in the go routine, when a encrypted message is written to 'item' the int channel, 'done' , gets 1 written to it. 
each time done gets a 1 written to it, it is read off in the main function and added to the 'sync'. when the integer 'sync' is greater than or equal to the length of the original message array, i.e. every message has been encrypted and printed, then the loop breaks and all the channels are closed.