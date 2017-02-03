{{ define "body" }}
<h1>Add Section</h1>
<form method="POST">
    <div class="form-group">
        <label for="Name">Name</label>
        <input type="text" class="form-control" name="Name" placeholder="Enter Section Name" value="{{ .Section.Name }}"/>
    </div>
    
    <div class="form-group">
        <label for="Icon">Icon Code</label>
        <input type="text" class="form-control" name="Icon" placeholder="Enter Section Icon code" value="{{ .Section.Icon }}"/>
    </div>
    <div class="form-group">
        <label for="Tagline">Tagline</label>
        <input type="text" class="form-control" name="Tagline" placeholder="Enter Section Tagline" value="{{ .Section.Tagline}}"/>
    </div>
    <input type="hidden" name="Id" value="{{ .Section.Id }}">
    <div class="form-group">
        <label for="Content">Content</label>
        <textarea class="form-control" name="Content" />{{ .Section.Content}}</textarea>
    </div>
    <button class="btn btn-primary" >Add</button>
</form>

{{ end }}
