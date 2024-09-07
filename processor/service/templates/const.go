package templates

const TotalBalanceTemplate = `

<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>User Summary</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f9;
            margin: 0;
            padding: 20px;
        }

        .container {
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            max-width: 600px;
            margin: 0 auto;
        }

        h2 {
            color: #2c3e50;
            text-align: center;
        }

        table {
            width: 100%;
            border-collapse: collapse; /* Ensure borders collapse */
            margin: 20px auto; /* Center table horizontally */
        }

        th,
        td {
            padding: 12px;
            text-align: center; /* Center text */
            border: 1px solid #ddd; /* Add border to cells */
        }

        th {
            background-color: #2ecc71; /* Green color */
            color: white;
            padding: 15px; /* Add padding to the header cells */
        }

        td {
            background-color: #ecf0f1;
            padding: 15px; /* Add padding to the data cells */
        }

        .numeric {
            font-family: 'Courier New', Courier, monospace; /* Monospace font for numeric alignment */
            width: 150px; /* Fixed width for numeric cells */
            text-align: right; /* Right align numeric values */
        }

        .balance {
            font-size: 1.5em;
            color: #27ae60; /* Darker green */
        }

        .credit,
        .debit {
            color: #e74c3c;
        }

        .footer {
            margin-top: 20px;
            text-align: center;
            font-size: 0.9em;
            color: #7f8c8d;
        }

        .logo {
            margin-top: 20px;
            text-align: center;
        }
    </style>
</head>

<body>
    <div class="container">
        <h2>Hello, {{.Contact.Name}}!</h2>
        <p>Here is your latest transaction summary:</p>

        <table>
            <tr>
                <th>Total Balance</th>
                <td class="numeric balance">${{printf "%.2f" .TotalBalance}}</td>
            </tr>
            <tr>
                <th>Average Credit</th>
                <td class="numeric credit">${{printf "%.2f" .AvgCredit}}</td>
            </tr>
            <tr>
                <th>Average Debit</th>
                <td class="numeric debit">${{printf "%.2f" .AvgDebit}}</td>
            </tr>
        </table>

        <h3 style="text-align: center;">Transactions per Month:</h3>
        <table>
            <tr>
                <th>Month</th>
                <th>Transaction Count</th>
            </tr>
            {{range $month, $count := .TransactionsPerMonth}}
            <tr>
                <td>{{$month}}</td>
                <td class="numeric">{{$count}}</td>
            </tr>
            {{end}}
        </table>

        <div class="footer">
            <p>Thank you for using our service!</p>
            <p>Contact us at support@stori.com</p>
            <div class="logo">
                <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/Stori_Logo_2023.svg/1920px-Stori_Logo_2023.svg.png" alt="Company Logo" width="150">
            </div>
        </div>
    </div>
</body>

</html>



`
