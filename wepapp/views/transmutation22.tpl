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
        <pre>
            N1(A1, Z1) + N2(A2, Z2) -----> N3(A3, Z3) + N4(A4, Z4)
            A1 + A2 = A3 + A4 ,    Z1 + Z2 = Z3 + Z4
        </pre>
    </div>
    <nav>
        <ul>
            <li><a href="/isotopes">Table of Stable Isotopes</a></li>
            <li><a href="/transmutations22">Table of conversion 2 Nucleides into 2 Nucleides</a></li>
            <li><a href="/fusionsfission">Tables of Fusion and Fission</a></li>
            <li>Tables for В,С, N, O, Al, Ni, Cu, W</li>
        </ul>
    </nav>
</header>

<div id="stable">
    <form action="/transmutations22" method="post"  class="form form-inline">
        <strong>search by:</strong>

        <p class="form-group">
            <label for="element">Element1</label>
            <input class="form-control" type="text" name="element1"  id="element1" size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="a">A1</label>
            <input class="form-control" type="text" name="A1"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="z">Z1</label>
            <input class="form-control" type="text" name="Z1"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="element">Element2</label>
            <input class="form-control" type="text" name="element2" id="element2"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="a">A2</label>
            <input class="form-control" type="text" name="A2"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="z">Z2</label>
            <input class="form-control" type="text" name="Z2"  size="2" value=""/>
        </p>


        <p class="form-group">
            <label for="element">Element3</label>
            <input class="form-control" type="text" name="element3"  id="element3"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="a">A3</label>
            <input class="form-control" type="text" name="A3"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="z">Z3</label>
            <input class="form-control" type="text" name="Z3"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="element">Element4</label>
            <input class="form-control" type="text" name="element4" id="element4"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="a">A4</label>
            <input class="form-control" type="text" name="A4"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="z">Z4</label>
            <input class="form-control" type="text" name="Z4"  size="2" value=""/>
        </p>

        <input type="submit" name="Filter" value="filter"/>
    </form>
{{if gt .paginator.PageNums 1}}
    <ul class="pagination pagination-sm">
    {{if .paginator.HasPrev}}
        <li><a href="{{.paginator.PageLinkFirst}}">First Page</a></li>
        <li><a href="{{.paginator.PageLinkPrev}}">&lt;</a></li>
    {{else}}
        <li class="disabled"><a>First Page</a></li>
        <li class="disabled"><a>&lt;</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <li{{if $.paginator.IsActive .}} class="active"{{end}}>
            <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </li>
    {{end}}
    {{if .paginator.HasNext}}
        <li><a href="{{.paginator.PageLinkNext}}">&gt;</a></li>
        <li><a href="{{.paginator.PageLinkLast}}">Last Page</a></li>
    {{else}}
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>Last Page</a></li>
    {{end}}
    </ul>
{{end}}
    <table class="table table-striped m900">
        <thead>
        <tr>
            <th>Element1</th>
            <th>A1</th>
            <th>Z2</th>
            <th>Element2</th>
            <th>A2</th>
            <th>Z2</th>
            <th>Element3</th>
            <th>A3</th>
            <th>Z3</th>
            <th>Element4</th>
            <th>A4</th>
            <th>Z4</th>
            <th>MeV</th>
        </tr>
        </thead>
        <tbody>
        {{range $transmutation := .Transmutations }}
        <tr>
            <td>{{$transmutation.Element1}}</td>
            <td>{{$transmutation.A1}}</td>
            <td>{{$transmutation.Z1}}</td>
            <td>{{$transmutation.Element2}}</td>
            <td>{{$transmutation.A2}}</td>
            <td>{{$transmutation.Z2}}</td>
            <td>{{$transmutation.Element3}}</td>
            <td>{{$transmutation.A3}}</td>
            <td>{{$transmutation.Z3}}</td>
            <td>{{$transmutation.Element4}}</td>
            <td>{{$transmutation.A4}}</td>
            <td>{{$transmutation.Z4}}</td>
            <td>{{$transmutation.Mev}}</td>
        </tr>
        {{ end }}
        </tbody>
    </table>
{{if gt .paginator.PageNums 1}}
    <ul class="pagination pagination-sm">
    {{if .paginator.HasPrev}}
        <li><a href="{{.paginator.PageLinkFirst}}">First Page</a></li>
        <li><a href="{{.paginator.PageLinkPrev}}">&lt;</a></li>
    {{else}}
        <li class="disabled"><a>First Page</a></li>
        <li class="disabled"><a>&lt;</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <li{{if $.paginator.IsActive .}} class="active"{{end}}>
            <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </li>
    {{end}}
    {{if .paginator.HasNext}}
        <li><a href="{{.paginator.PageLinkNext}}">&gt;</a></li>
        <li><a href="{{.paginator.PageLinkLast}}">Last Page</a></li>
    {{else}}
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>Last Page</a></li>
    {{end}}
    </ul>
{{end}}
</div>
<footer>
    <div class="author">
        Official website:
        <a href="http://www.quantumheat.org/index.php/en/">QuantumHeat.org</a>
    </div>
</footer>

<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
<script src="/static/js/jquery-3.3.1.min.js"></script>
<script src="/static/js/jquery.autocomplete.min.js"></script>
<script>
    $('#element1').autocomplete({
        serviceUrl: '/transmutations22/element1'
    });
    $('#element2').autocomplete({
        serviceUrl: '/transmutations22/element2'
    });
    $('#element3').autocomplete({
        serviceUrl: '/transmutations22/element3'
    });
    $('#element4').autocomplete({
        serviceUrl: '/transmutations22/element4'
    });
</script>
</body>
</html>


