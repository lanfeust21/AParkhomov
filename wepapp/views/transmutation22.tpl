<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>Table of Stable Transmutations - data provided by Alexander Parkhomov</title>

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
        .source {
            color: #3399FF;
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
    <h1 class="logo">2 to 2 atom nucleon transfer reactions- data provided by Alexander Parkhomov</h1>
    <h2>Facilitated by the Martin Fleischmann Memorial Project @ QuantumHeat.org</h2>

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
    <h2>2 to 2 atom nucleon transfer reactions</h2>
    <div class="description">
        <pre>
            <span class="source">first pair (sources)</span>     -----> second pair (products)
            <span class="source">N1(A1, Z1) + N2(A2, Z2)</span> -----> N3(A3, Z3) + N4(A4, Z4)
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
        <p class="form-group">
            <label for="sortorder">Sort Order</label>
            <select class="form-control"  name="sortorder">
                <option value="-Mev" {{if eq .Sortorder "-Mev"}} selected="true"{{end}}>Mev desc</option>
                <option value="-A1,-Z1,-Mev" {{if eq .Sortorder "-A1,-Z1,-Mev"}} selected="true"{{end}}>A1 desc,Z1 desc,Mev desc</option>
                <option value="-A2,-Z2,-Mev" {{if eq .Sortorder "-A2,-Z2,-Mev"}} selected="true"{{end}}>A2 desc,Z2 desc,Mev desc</option>
                <option value="-A3,-Z3,-Mev" {{if eq .Sortorder "-A3,-Z3,-Mev"}} selected="true"{{end}}>A3 desc,Z3 desc,Mev desc</option>
                <option value="-A4,-Z4,-Mev" {{if eq .Sortorder "-A4,-Z4,-Mev"}} selected="true"{{end}}>A4 desc,Z4 desc,Mev desc</option>
            </select>
        </p>
        <input type="submit" name="Filter" value="filter"/>
    </form>
    {{ $QS := .QueryString }}
    <ul class="pagination pagination-sm">
    {{if gt .paginator.PageNums 1}}
    {{if .paginator.HasPrev}}
        <li>{{ Link .paginator.PageLinkFirst "First Page" $QS }}</li>
        <li>{{ Link .paginator.PageLinkPrev "&lt;" $QS }}</li>
    {{else}}
        <li class="disabled"><a>First Page</a></li>
        <li class="disabled"><a>&lt;</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <li{{if $.paginator.IsActive .}} class="active"{{end}}>
        {{ $pagelink :=  $.paginator.PageLink $page }}
        {{ Link $pagelink $page $QS }}
        </li>
    {{end}}
    {{if .paginator.HasNext}}
        <li>{{ Link .paginator.PageLinkNext "&gt;" $QS }}</li>
        <li>{{ Link .paginator.PageLinkLast "Last Page" $QS }}</li>
    {{else}}
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>Last Page</a></li>
    {{end}}
    {{end}}
        <li>{{ Link "/transmutations22" "Direct Link" $QS }}</li>
        <li>{{ Link "/transmutations22/csv" "CSV Export" $QS }}</li>
        <li>Count:<strong>{{.Count}}</strong></li>
    </ul>
    <table class="table table-striped m900">
        <thead>
        <tr>
            <th class="source">Element1</th>
            <th class="source">A1</th>
            <th class="source">Z2</th>
            <th class="source">Element2</th>
            <th class="source">A2</th>
            <th class="source">Z2</th>
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
            <td class="source">{{$transmutation.Element1}}</td>
            <td class="source">{{$transmutation.A1}}</td>
            <td class="source">{{$transmutation.Z1}}</td>
            <td class="source">{{$transmutation.Element2}}</td>
            <td class="source">{{$transmutation.A2}}</td>
            <td class="source">{{$transmutation.Z2}}</td>
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
    <ul class="pagination pagination-sm">
    {{if gt .paginator.PageNums 1}}
    {{if .paginator.HasPrev}}
        <li>{{ Link .paginator.PageLinkFirst "First Page" $QS }}</li>
        <li>{{ Link .paginator.PageLinkPrev "&lt;" $QS }}</li>
    {{else}}
        <li class="disabled"><a>First Page</a></li>
        <li class="disabled"><a>&lt;</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <li{{if $.paginator.IsActive .}} class="active"{{end}}>
        {{ $pagelink :=  $.paginator.PageLink $page }}
        {{ Link $pagelink $page $QS }}
        </li>
    {{end}}
    {{if .paginator.HasNext}}
        <li>{{ Link .paginator.PageLinkNext "&gt;" $QS }}</li>
        <li>{{ Link .paginator.PageLinkLast "Last Page" $QS }}</li>
    {{else}}
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>Last Page</a></li>
    {{end}}
    {{end}}
        <li>{{ Link "/transmutations22" "Direct Link" $QS }}</li>
        <li>{{ Link "/transmutations22/csv" "CSV Export" $QS }}</li>
        <li>Count:<strong>{{.Count}}</strong></li>
    </ul>
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


