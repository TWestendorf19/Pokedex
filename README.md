# Pokedex

  This project is based on a Pokedex project guide that can be found on boot.dev, I highly recommend checking them out if you want to work on backend development.  

  At this stage it is a simple command line interface that has a help and detail command to assist the user, a few traversal commands, and an exit command to end the program.

  ## Commands  
    
  * Detail returns the details of a specific command and is used with the format 'detail 'command''.  

  * Help returns the details of all available commands and is used with the format 'help' or 'help nameonly'. (Adding the nameonly special command just lists the names of all commands with no descriptions)  

  * Exit ends the program and is used with the format(Special commands added after exit will be printed back out to the user upon exit)  

  * Map returns a list of all the "areas" 20 at a time and increments through all of them with each call. (adding the name of a location e.g. canalave-city or mt-coronet will return the areas within that location).  

  * Mapb is similar to the map command but iterates backwards through the list. (adding the name of a location works the same as map command)  

  * Locations will return a list of the "locations" in the game. These locations contain at least one "area".  

### ToDo:  
  
- Add Locations command to display possible locations to enter in map commands
- Add Caching 
- Add Explore
- Add Catching
- Add Inspect
- Add Pokedex command

### Potential Feature Additions:  
  
- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
