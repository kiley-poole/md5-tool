# Marvel API MD5 Tool

This is a quick Go tool to generate the hash needed for the API validation on Marvel API calls.  
## The URL parameters needed for auth are:  
`MARVELURL?ts={timestamp}&apikey={publickey}&hash={hash}`  

## How to use  
This tool generates a random int, requests your private and public keys and then creates the proper MD5 hash needed for access. Just copy the returned url parameter string and append it to any Marvel gateway request for accesss.

## Why was this created?  
I made this is a quick way to test to out the Marvel API in browser without having to rely on their swagger documentation.
