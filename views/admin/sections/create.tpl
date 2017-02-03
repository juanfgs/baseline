{{ define "body" }}
<h1>Add Section</h1>
<form method="POST">
    <div class="form-group">
        <label for="Name">Name</label>
        <input type="text" class="form-control" name="Name" placeholder="Enter Section Name"/>
    </div>
    
    <div class="form-group">
        <label for="Icon">Icon Code</label>
        <input type="text" class="form-control" name="Icon" placeholder="Enter Section Icon code"/>
    </div>
    <div class="form-group">
        <label for="Tagline">Tagline</label>
        <input type="text" class="form-control" name="Tagline" placeholder="Enter Section Tagline"/>
    </div>
    <div class="form-group">
        <label for="Content">Content</label>
        <textarea class="form-control" name="Content" /></textarea>
    </div>
    <button class="btn btn-primary" >Add</button>
</form>

{{ end }}
