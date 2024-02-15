import java.util.Scanner;
import java.util.HashMap;
import java.util.Map;
import java.io.IOException;
import org.json.JSONObject;
import org.json.JSONException;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class CurrencyConverter {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Map<String, Double> cache = new HashMap<>();

        System.out.print("Enter your currency code: ");
        String x = scanner.nextLine().toUpperCase();

        OkHttpClient client = new OkHttpClient();
        Request request = new Request.Builder()
                .url("http://www.floatrates.com/daily/" + x + ".json")
                .build();

        try {
            Response response = client.newCall(request).execute();
            String jsonData = response.body().string();
            JSONObject jsonObject = new JSONObject(jsonData);

            if (!x.equals("USD")) {
                cache.put("USD", jsonObject.getJSONObject("usd").getDouble("rate"));
            }
            if (!x.equals("EUR")) {
                cache.put("EUR", jsonObject.getJSONObject("eur").getDouble("rate"));
            }

            while (true) {
                System.out.print("Enter exchange currency: ");
                String y = scanner.nextLine().toUpperCase();

                if (y.isEmpty()) {
                    break;
                }

                System.out.print("Enter amount of " + x + ": ");
                String numStr = scanner.nextLine();
                double num = Double.parseDouble(numStr);

                System.out.println("Checking the cache...");
                if (cache.containsKey(y)) {
                    System.out.println("Oh! It is in the cache!");
                } else {
                    System.out.println("Sorry, but it is not in the cache!");
                    cache.put(y, jsonObject.getJSONObject(y.toLowerCase()).getDouble("rate"));
                }
                System.out.println("You received " + String.format("%.2f", (cache.get(y) * num)) + " " + y);
            }
        } catch (IOException | JSONException e) {
            e.printStackTrace();
        }
        scanner.close();
    }
}
