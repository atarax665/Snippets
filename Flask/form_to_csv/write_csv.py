import pandas as pd 
 
def add_data(name,addID,branch):
    isThere = False
    df1 = pd.read_csv('./record.csv')
    for index, row in df1.iterrows():
        if str(row['Addmission_No']) == str(addID):
            isThere = True
            status = "Student already registered. Please recheck Name and Addmission Number."
            return status

    if isThere != True: 
        df2 = pd.DataFrame({"Name":[name], 
        "Addmission_No":[addID], 
        "Branch":[branch]}) 
        dff = df1.append(df2, ignore_index = True)
        dff.to_csv(r'./record.csv', index = False)
        status = "Registration Successful"
        return status