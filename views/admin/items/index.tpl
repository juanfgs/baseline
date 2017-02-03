{{ define "body"}}
<h1>Items</h1>
<div class="panel panel-default">
  <div class="panel-body">
<table class="table table-striped">
    <tr>
        <th>Id</th>
        <th>Name</th>
        <th>Value</th>
        <th>Actions</th>
    </tr>
    {{range $key, $item := .ItemList}}
    <tr>
        <td>{{ $item.Id}}</td>
        <td>{{ $item.Name}}</td>
        <td>{{ $item.Value}}</td>
        <td></td>
    </tr>
</table>    
    {{end }}

  </div>
</div>

{{ end }}
