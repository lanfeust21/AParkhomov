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
    <h1 class="logo">Table of Stable Transmutations - Alexander Parkhomov Nuclear Data</h1>
    <h2>Martin Fleischmann Memorial Project QuantumHeat.org</h2>

    <nav>
        <ul>
            <li><a href="/isotopes">Table of Stable Isotopes</a></li>
            <li><a href="/transmutations22">Table of conversion 2 Nucleides into 2 Nucleides</a></li>
            <li><a href="/fusions">Tables of Fusions</a></li>
            <li><a href="/fissions">Tables of Fissions</a></li>
        </ul>
    </nav>
</header>

<div id="stable">
    <h2>Stable Transmutations</h2>
    <div class="description">
        <pre>
            first pair              -----> second pair
            N1(A1, Z1) + N2(A2, Z2) -----> N3(A3, Z3) + N4(A4, Z4)
            A1 + A2 = A3 + A4 ,    Z1 + Z2 = Z3 + Z4
        </pre>
    </div>
    <form action="/transmutations22" method="post"  class="form form-inline" id="transmutations">
        <strong>search by:</strong>

        <p class="form-group">
            <label for="element">Element1</label>
            <input class="form-control" type="text" name="element1"  id="element1" size="2" value="{{.Element1}}"/>
        </p>
        <p class="form-group">
            <label for="a">A1</label>
            <input class="form-control" type="text" name="A1"  size="2" value="{{.A1}}"/>
        </p>
        <p class="form-group">
            <label for="z">Z1</label>
            <input class="form-control" type="text" name="Z1"  size="2" value="{{.Z1}}"/>
        </p>
        <p class="form-group">
            <label for="element">Element2</label>
            <input class="form-control" type="text" name="element2" id="element2"  size="2" value="{{.Element2}}"/>
        </p>
        <p class="form-group">
            <label for="a">A2</label>
            <input class="form-control" type="text" name="A2"  size="2" value="{{.A2}}"/>
        </p>
        <p class="form-group">
            <label for="z">Z2</label>
            <input class="form-control" type="text" name="Z2"  size="2" value="{{.Z2}}"/>
        </p>
        &nbsp;--&gt;&nbsp;
        <p class="form-group">
            <label for="element">Element3</label>
            <input class="form-control" type="text" name="element3"  id="element3"  size="2" value="{{.Element3}}"/>
        </p>
        <p class="form-group">
            <label for="a">A3</label>
            <input class="form-control" type="text" name="A3"  size="2" value="{{.A3}}"/>
        </p>
        <p class="form-group">
            <label for="z">Z3</label>
            <input class="form-control" type="text" name="Z3"  size="2" value="{{.Z3}}"/>
        </p>
        <p class="form-group">
            <label for="element">Element4</label>
            <input class="form-control" type="text" name="element4" id="element4"  size="2" value="{{.Element4}}"/>
        </p>
        <p class="form-group">
            <label for="a">A4</label>
            <input class="form-control" type="text" name="A4"  size="2" value="{{.A4}}"/>
        </p>
        <p class="form-group">
            <label for="z">Z4</label>
            <input class="form-control" type="text" name="Z4"  size="2" value="{{.Z4}}"/>
        </p>
        <p class="form-group">
            <label for="limit">Result per page</label>
            <select class="form-control"  name="limit">
                <option value="50" {{if eq .Limit "50"}} selected="true"{{end}}>50</option>
                <option value="100" {{if eq .Limit "100"}} selected="true"{{end}}>100</option>
                <option value="250" {{if eq .Limit "250"}} selected="true"{{end}}>250</option>
                <option value="500" {{if eq .Limit "500"}} selected="true"{{end}}>500</option>
                <option value="1000" {{if eq .Limit "1000"}} selected="true"{{end}}>1000</option>
            </select>
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
        &nbsp;
        Source:
        <a href="https://github.com/lanfeust21/AParkhomov">https://github.com/lanfeust21/AParkhomov</a>
    </div>
</footer>

<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
<script src="/static/js/jquery-3.3.1.min.js"></script>
<script src="/static/js/jquery.autocomplete.min.js"></script>
<script>
    $('#element1').autocomplete({
        serviceUrl: '/transmutations22/element1',
        select: function(event, ui) { $("#transmutations").submit(); }
    });
    $('#element2').autocomplete({
        serviceUrl: '/transmutations22/element2',
        select: function(event, ui) { $("#transmutations").submit(); }
    });
    $('#element3').autocomplete({
        serviceUrl: '/transmutations22/element3',
        select: function(event, ui) { $("#transmutations").submit(); }
    });
    $('#element4').autocomplete({
        serviceUrl: '/transmutations22/element4',
        select: function(event, ui) { $("#transmutations").submit(); }
    });
</script>
</body>
</html>


