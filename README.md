# Inventory Aging

Inventory aging, and product tracking are two important issues that are usually monitoring through `Batch Control`.
As some system not having batch control setup, doing the above will take lots of manual efforts depending of the volume of data to be processed. specially once you are working with multiple warehouses.

Th eonly input data required is the transactions recorded in the system exported as `csv` file, and the program will generate 2 csv files that can be opened easily in MS excell or any other spreed sheet, these 2 files are:

1. Inventory aging shouing remaining quantity of each reciept, with the date and year of reciept.
2. Inventory distribution showing when each quantity of each batch/reciept had been issed. 

I made this repository to same such time, by programming the process using GO language, and making it availabe for public use hoping some people to get benifit of it (bytuning the code to match their setup).
