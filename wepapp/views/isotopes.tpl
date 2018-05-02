<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>Welcome to Alexander Parkhomov Nuclear Data</title>

    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/bootstrap-table.min.css" rel="stylesheet">
    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
    <style>
        .m900 {
            width: 900px;
        }
        .m300 {
            width: 900px;
        }
    </style>
</head>
<body>
<header>
    <h1 class="logo">Table of Stable Isotopes - Alexander Parkhomov Nuclear Data</h1>
    <h2>Martin Fleischmann Memorial Project QuantumHeat.org</h2>
    <div class="description">
        All stable isotopes and radioactive 1n 60Со 90Sr 137Cs 232Th 235U 238U 239Pu
    </div>
    <nav>
        <ul>
            <li><a href="/isotopes">Table of Stable Isotopes</a></li>
            <li>Table of conversion 2 Nucleides into 2 Nucleides</li>
            <li>Tables of Fusion and Fission</li>
            <li>Tables for В,С, N, O, Al, Ni, Cu, W</li>
        </ul>
    </nav>
</header>
<div id="stable">
    <form action="/isotopes" method="post"  class="form m300">
        <strong>search by:</strong>
        <p class="form-group">
            <label for="element">Element</label>
            <input class="form-control" type="text" name="element" value=""/>
        </p>
        <p class="form-group">
            <label for="a">A</label>
            <input class="form-control" type="text" name="A" value=""/>
        </p>
        <p class="form-group">
            <label for="z">Z</label>
            <input class="form-control" type="text" name="Z" value=""/>
        </p>
        <input type="submit" name="Filter" value="filter"/>
    </form>

    <table class="table table-striped m900">
        <thead>
        <tr>
            <th>element</th>
            <th>A</th>
            <th>Z</th>
            <th>a.e.m</th>
            <th>MeV</th>
            <th>%</th>
        </tr>
        </thead>
        <tbody>
        {{range $isotopes := .Isotopes }}
        <tr>
            <td>{{$isotopes.Element}}</td>
            <td>{{$isotopes.A}}</td>
            <td>{{$isotopes.Z}}</td>
            <td>{{$isotopes.AEM}}</td>
            <td>{{$isotopes.Mev}}</td>
            <td>{{$isotopes.Percent}}</td>
        </tr>
        {{ end }}
        </tbody>
    </table>

</div>
<footer>
    <div class="author">
        Official website:
        <a href="http://www.quantumheat.org/index.php/en/">QuantumHeat.org</a>
    </div>
</footer>

<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
<script src="/static/js/jquery-3.3.1.min.js"></script>
<!-- Include all compiled plugins (below), or include individual files as needed -->
<script src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/bootstrap-table.min.js"></script>
</body>
</html>