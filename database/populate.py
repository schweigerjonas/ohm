import sqlite3
import pandas as pd

# create connection to database
con = sqlite3.connect("budget.db")
c = con.cursor()

# get categories from file
df = pd.read_csv("./categories.csv")
data = list(df.itertuples(index=False, name=None))

# insert categories
c.executemany(
    "INSERT INTO categories (type, category, subcategory) values (?, ?, ?)", data
)

# close connection
con.commit()
con.close()
