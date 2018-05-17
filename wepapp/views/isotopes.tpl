<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>Table of Stable Isotopes - data provided by Alexander Parkhomov</title>

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
        .autocomplete-suggestions { border: 1px solid #999; background: #FFF; overflow: auto; }
        .autocomplete-suggestion { padding: 2px 5px; white-space: nowrap; overflow: hidden; }
        .autocomplete-selected { background: #F0F0F0; }
        .autocomplete-suggestions strong { font-weight: normal; color: #3399FF; }
        .autocomplete-group { padding: 2px 5px; }
        .autocomplete-group strong { display: block; border-bottom: 1px solid #000; }
        .ui-autocomplete-custom {
            background: #87ceeb;
            z-index: 2;
        }
    </style>
</head>
<body>
<header>
    <h1 class="logo">Considered Isotopes - data provided by Alexander Parkhomov</h1>
    <h2>Facilitated by the Martin Fleischmann Memorial Project @ QuantumHeat.org</h2>
    <div class="description">
        All stable isotopes and radioactive 1n 60Со 90Sr 137Cs 232Th 235U 238U 239Pu
    </div>
    <nav>
        <ul>
            <li><a href="/isotopes">Considered Isotopes</a></li>
            <li><a href="/transmutations22">2 to 2 atom nucleon transfer reactions</a></li>
            <li><a href="/fusions">Fusion reactions</a></li>
            <li><a href="/fissions">Fission reactions</a></li>
        </ul>
    </nav>
</header>
<div id="stable">
    <form action="/isotopes" method="post"  class="form m300 form-inline" id="isotopes">
        <strong>search by:</strong>
        <p class="form-group">
            <label for="element">Element</label>
            <input class="form-control" type="text" name="element" size="2" id="Element" value="{{.Element}}"/>
        </p>
        <p class="form-group">
            <label for="a">A</label>
            <input class="form-control" type="text" name="A" value="{{.A}}"/>
        </p>
        <p class="form-group">
            <label for="z">Z</label>
            <input class="form-control" type="text" name="Z" value="{{.Z}}"/>
        </p>
        <input type="submit" name="Filter" value="filter"/>
    </form>
    <div>count:{{.Count}}</div>
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
        &nbsp;
        Source:
        <a href="https://github.com/lanfeust21/AParkhomov">https://github.com/lanfeust21/AParkhomov</a>
    </div>
</footer>

<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
<script src="/static/js/jquery-3.3.1.min.js"></script>
<script src="/static/js/jquery.autocomplete.min.js"></script>
<script>
    $('#Element').autocomplete({
        serviceUrl: '/isotopes/element',
        select: function(event, ui) { $("#isotopes").submit(); }
    });
</script>
</body>
</html>