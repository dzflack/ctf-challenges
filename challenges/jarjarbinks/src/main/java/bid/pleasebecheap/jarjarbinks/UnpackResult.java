package bid.pleasebecheap.jarjarbinks;

import java.util.ArrayList;
import java.util.List;
import java.io.File;
import java.util.concurrent.atomic.AtomicInteger;

public class UnpackResult {
    private final String description;
    private final ArrayList<Result> contents;
    private final String errors;

    public UnpackResult(List<String> contents) {
        this.description = "Jar downloaded and extracted in /tmp/";
        this.errors = "";

        this.contents = new ArrayList<Result>();
        AtomicInteger count = new AtomicInteger(0);

        for (String s : contents){
            Result thing = new Result(Integer.toString(count.incrementAndGet()), s);
            this.contents.add(thing);
        }
    }

    public UnpackResult(Exception e) {
        File location = new File(HelloController.class.getProtectionDomain().getCodeSource().getLocation().getPath());
        String x = location.toString().split("classes")[0];

        this.errors = x + ":  " + e.toString();
        this.description = "";
        this.contents = new ArrayList<Result>();
    }

    public ArrayList<Result> getContents() {
        return this.contents;
    }

    public String getDescription(){
        return this.description;
    }

    public String getErrors(){
        return this.errors;
    }

}