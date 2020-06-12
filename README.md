# Inventory Aging

#inventory #aging, and product #tracking are two important issues that are usually monitored through #batchcontrol. As some systems not having batch control setup, doing the above (assuming #fifo is in place) will take lots of manual efforts depending of the volume of data to be processed. specially once you are working with multiple warehouses.

I made this repository to save calculations time, by programming the process using #go language, and making it available for public use hoping some people to get benefit of it (by tuning the code to match their setup).

The only input data required is the transactions record exported from the system as csv file, and the program will generate 2 csv files that can be opened easily in MS excel or any other spread sheet, these 2 files are:

1. Inventory aging showing remaining quantity of each receipt, with the date and year of receipt.

2. Inventory distribution showing when each quantity of each receipt had been issued.
