<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>Table of Nuclear fission - data provided by Alexander Parkhomov</title>

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
    <h1 class="logo">Table of Nuclear fission - data provided by Alexander Parkhomov</h1>
    <h2>Facilitated by the Martin Fleischmann Memorial Project @ QuantumHeat.org</h2>

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
    <h2>Nuclear fission</h2>
    <div class="description">
        <pre>
            <span class="source">source</span>  -----> pair (products)
            <span class="source">N(A, Z)</span> -----> N1(A1, Z1) + N2(A2, Z2)
                A = A1 + A2 ,  Z = Z1 + Z2
        </pre>
    </div>
    <form action="/fissions" method="post"  class="form form-inline">
        <strong>search by:</strong>
        <p class="form-group">
            <label for="element">Element</label>
            <input class="form-control" type="text" name="element"  id="element"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="a">A</label>
            <input class="form-control" type="text" name="A"  size="2" value=""/>
        </p>
        <p class="form-group">
            <label for="z">Z</label>
            <input class="form-control" type="text" name="Z"  size="2" value=""/>
        </p>
        &nbsp;--&gt;&nbsp;
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
        <li>{{ Link "/fissions" "Direct Link" $QS }}</li>
        <li>{{ Link "/fissions/csv" "CSV Export" $QS }}</li>
        <li>Count:<strong>{{.Count}}</strong></li>
    </ul>
    <table class="table table-striped m900">
        <thead>
        <tr>
            <th class="source">Element</th>
            <th class="source">A</th>
            <th class="source">Z</th>
            <th>Element1</th>
            <th>A1</th>
            <th>Z2</th>
            <th>Element2</th>
            <th>A2</th>
            <th>Z2</th>

            <th>MeV</th>
        </tr>
        </thead>
        <tbody>
        {{range $fission := .Fissions }}
        <tr>
            <td class="source">{{$fission.Element}}</td>
            <td class="source">{{$fission.A}}</td>
            <td class="source">{{$fission.Z}}</td>
            <td>{{$fission.Element1}}</td>
            <td>{{$fission.A1}}</td>
            <td>{{$fission.Z1}}</td>
            <td>{{$fission.Element2}}</td>
            <td>{{$fission.A2}}</td>
            <td>{{$fission.Z2}}</td>
            <td>{{$fission.Mev}}</td>
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
        <li>{{ Link "/fissions" "Direct Link" $QS }}</li>
        <li>{{ Link "/fissions/csv" "CSV Export" $QS }}</li>
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
        serviceUrl: '/fissions/element1'
    });
    $('#element2').autocomplete({
        serviceUrl: '/fissions/element2'
    });

    $('#element').autocomplete({
        serviceUrl: '/fissions/element'
    });
</script>
</body>
</html>


