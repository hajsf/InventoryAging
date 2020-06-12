# Inventory Aging

Inventory aging, and product tracking are two important issues that are usually monitored through Batch Control. As some systems not having batch control setup, doing the above (assuming FIFO is in place) will take lots of manual efforts depending of the volume of data to be processed. specially once you are working with multiple warehouses.

The only input data required is the transactions record exported from the system as csv file, and the program will generate 2 csv files that can be opened easily in MS excel or any other spread sheet, these 2 files are:

Inventory aging showing remaining quantity of each receipt, with the date and year of receipt.

Inventory distribution showing when each quantity of each receipt had been issued.

I made this repository to save calculations time, by programming the process using GO language, and making it available for public use hoping some people to get benefit of it (by tuning the code to match their setup).
