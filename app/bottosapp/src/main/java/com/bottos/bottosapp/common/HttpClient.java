package com.bottos.bottosapp.common;

import org.apache.http.NoHttpResponseException;
import org.apache.http.client.HttpRequestRetryHandler;
import org.apache.http.client.config.RequestConfig;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.config.Registry;
import org.apache.http.config.RegistryBuilder;
import org.apache.http.conn.ConnectTimeoutException;
import org.apache.http.conn.socket.ConnectionSocketFactory;
import org.apache.http.conn.socket.LayeredConnectionSocketFactory;
import org.apache.http.conn.socket.PlainConnectionSocketFactory;
import org.apache.http.conn.ssl.SSLConnectionSocketFactory;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.impl.conn.PoolingHttpClientConnectionManager;
import org.apache.http.message.BasicHeader;
import org.apache.http.protocol.HttpContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.net.ssl.SSLException;
import javax.net.ssl.SSLHandshakeException;
import java.io.*;
import java.net.UnknownHostException;
import java.util.Iterator;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;


public class HttpClient {

    private static Logger logger = LoggerFactory.getLogger(HttpClient.class);
    private static CloseableHttpClient httpClient;
    private static final Lock lock = new ReentrantLock();
    public static String maxTotal;
    public static String maxPerRoute;

    public HttpClient() {
    }

    public static CloseableHttpClient getCloseableHttpClient() {
        if (httpClient == null) {
            try {
                lock.lock();
                if (httpClient == null) {
                    httpClient = initCloseableHttpClient();
                }
            } finally {
                lock.unlock();
            }
        }

        return httpClient;
    }

    private static CloseableHttpClient initCloseableHttpClient() {
        ConnectionSocketFactory plainsf = PlainConnectionSocketFactory.getSocketFactory();
        LayeredConnectionSocketFactory sslsf = SSLConnectionSocketFactory.getSocketFactory();
        Registry socketFactoryRegistry = RegistryBuilder.create().register("http", plainsf).register("https", sslsf).build();

        PoolingHttpClientConnectionManager cm = new PoolingHttpClientConnectionManager(socketFactoryRegistry);
        if (maxTotal == null) {
            maxTotal = "200";
        }

        if (maxPerRoute == null) {
            maxPerRoute = "200";
        }

        cm.setMaxTotal(Integer.valueOf(maxTotal).intValue());
        cm.setDefaultMaxPerRoute(Integer.valueOf(maxPerRoute).intValue());
        HttpRequestRetryHandler httpRequestRetryHandler = new HttpRequestRetryHandler() {
            public boolean retryRequest(IOException exception, int executionCount, HttpContext context) {
                return executionCount >= 5 ? false : (exception instanceof NoHttpResponseException ? true : (exception instanceof SSLHandshakeException ? false : (exception instanceof InterruptedIOException ? false : (exception instanceof UnknownHostException ? false : (exception instanceof ConnectTimeoutException ? false : (exception instanceof SSLException ? false : false))))));
            }
        };
        CloseableHttpClient closeableHttpClient = HttpClients.custom().setConnectionManagerShared(true).setConnectionManager(cm).setRetryHandler(httpRequestRetryHandler).build();
        return closeableHttpClient;
    }

    public static String sendHttpPost(String url, String sendStr, Integer conncetTimeout, Integer socketTimeout, Map<String, String> headers) {
        RequestConfig requestConfig = RequestConfig.custom().setSocketTimeout(socketTimeout.intValue()).setConnectTimeout(conncetTimeout.intValue()).build();
        HttpPost httpPost = new HttpPost(url);
        httpPost.setConfig(requestConfig);
        StringEntity entity = new StringEntity(sendStr.toString(), "utf-8");
        String key;
        if (headers == null) {
            httpPost.setHeader("Content-Type", "application/json");
            entity.setContentEncoding(new BasicHeader("Content-Type", "application/json"));
        } else {
            Set<String> keys = headers.keySet();
            Iterator i = keys.iterator();

            while (i.hasNext()) {
                key = (String) i.next();
                httpPost.addHeader(key, (String) headers.get(key));
            }
        }

        httpPost.setEntity(entity);
        CloseableHttpClient closeableHttpClient = getCloseableHttpClient();
        StringBuilder strber = new StringBuilder();
        key = null;
        InputStream inStream = null;
        BufferedReader reader = null;

        String var14;
        try {
            logger.info("http postConnect，url:" + url + "  req：" + sendStr);
            CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpPost);
            inStream = httpResponse.getEntity().getContent();
            reader = new BufferedReader(new InputStreamReader(inStream, "utf-8"));
            String line = null;

            while ((line = reader.readLine()) != null) {
                strber.append(line);
            }

            logger.info("return：" + strber);
            var14 = strber.toString();
            return var14;
        } catch (Exception var24) {
            var24.printStackTrace();
            var14 = null;
        } finally {
            try {
                if (inStream != null) {
                    inStream.close();
                }

                if (reader != null) {
                    reader.close();
                }
            } catch (IOException var23) {
                var23.printStackTrace();
                return null;
            }

        }

        return var14;
    }

    public static String sendHttpGet(String url, String sendStr, Integer conncetTimeout, Integer socketTimeout, Map<String, String> headers) {
        RequestConfig requestConfig = RequestConfig.custom().setSocketTimeout(socketTimeout.intValue()).setConnectTimeout(conncetTimeout.intValue()).build();
        HttpGet httpGet = new HttpGet(url + "?" + sendStr);
        httpGet.setConfig(requestConfig);
        String key;
        if (headers != null) {
            Set<String> keys = headers.keySet();
            Iterator i = keys.iterator();

            while (i.hasNext()) {
                key = (String) i.next();
                httpGet.addHeader(key, (String) headers.get(key));
            }
        }

        CloseableHttpClient closeableHttpClient = getCloseableHttpClient();
        StringBuilder strber = new StringBuilder();
        key = null;
        InputStream inStream = null;
        BufferedReader reader = null;

        String var13;
        try {
            logger.info("http getConnect，url:" + url + "  req：" + sendStr);
            CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpGet);
            inStream = httpResponse.getEntity().getContent();
            reader = new BufferedReader(new InputStreamReader(inStream, "UTF-8"));
            String line = null;

            while ((line = reader.readLine()) != null) {
                strber.append(line);
            }

            logger.info("return：" + strber);
            var13 = strber.toString();
            return var13;
        } catch (Exception var23) {
            var23.printStackTrace();
            var13 = null;
        } finally {
            try {
                if (inStream != null) {
                    inStream.close();
                }

                if (reader != null) {
                    reader.close();
                }
            } catch (IOException var22) {
                var22.printStackTrace();
                return null;
            }

        }

        return var13;
    }

}
