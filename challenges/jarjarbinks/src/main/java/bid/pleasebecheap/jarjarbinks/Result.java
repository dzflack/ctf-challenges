package bid.pleasebecheap.jarjarbinks;

public class Result
{
    private String id;
    private String file;

    public Result(String id, String file){
        this.id = id;
        this.file = file;
    }

    public String getId(){
        return this.id;
    }
    
    public String getFile(){
        return this.file;
    }

}
	