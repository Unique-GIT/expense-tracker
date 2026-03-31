# expense-tracker
It is going to be an expense tracker with a simple excel based database and a wrapper on top of the cli for a whatsapp bot
Purpose: expense management from whatsapp itself, with minimal database and analyze features for tracking expense

It has the following functionality: 

1) Adding a Record:
    1. A Record is a tuple of (Name,Label,Cost)
    2. A Label is a fixed set ( enum ) to choose from, there will be a dropdown for the same

2) Add/Remove a Label: A Label is a category Name

3) Analyse: 
    1. Operation result gives us a table of ( Label, total cost, cost percentage)
        1. Total Cost : The total cost spent with that Label
        2. Cost Percentage: % of Total cost that was spent on that given label
    2. The Time duration has configuration: The total cost and cost percentage over this time duration
        1. Days 
        2. Weeks
        3. Months
        4. Years

