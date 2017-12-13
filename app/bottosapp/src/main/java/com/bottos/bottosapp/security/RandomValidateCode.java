package com.bottos.bottosapp.security;

import javax.imageio.ImageIO;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.awt.*;
import java.awt.image.BufferedImage;
import java.io.ByteArrayOutputStream;
import java.util.Random;


public class RandomValidateCode {
    private Random random = new Random();
    private String randString = "0123456789abcdefghijkmnpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";//Randomly generated string
    private int width = 80;// Picture wide
    private int height = 26;// Picture high
    private int lineSize = 40;// Interference line quantity
    private int stringNum = 4;// Random number of characters

    /**
     *
     * @return Get the font
     */
    private Font getFont() {
        return new Font("Fixedsys", Font.CENTER_BASELINE, 18);
    }

    /**
     *
     * @param fc
     * @param bc
     * @return Get color
     */
    private Color getRandColor(int fc, int bc) {
        if (fc > 255)
            fc = 255;
        if (bc > 255)
            bc = 255;
        int r = fc + random.nextInt(bc - fc - 16);
        int g = fc + random.nextInt(bc - fc - 14);
        int b = fc + random.nextInt(bc - fc - 18);
        return new Color(r, g, b);
    }

    /**
     * @param g
     * @param randomString
     * @param i
     * @return Get a random string
     */
    private String drowString(Graphics g, String randomString, int i) {
        g.setFont(getFont());
        g.setColor(new Color(random.nextInt(101), random.nextInt(111), random
                .nextInt(121)));
        String rand = String.valueOf(getRandomString(random.nextInt(randString
                .length())));
        randomString += rand;
        g.translate(random.nextInt(3), random.nextInt(3));
        g.drawString(rand, 13 * i, 16);
        return randomString;
    }

    /**
     *Draw interference lines
     * @param g
     */
    private void drowLine(Graphics g) {
        int x = random.nextInt(width);
        int y = random.nextInt(height);
        int xl = random.nextInt(13);
        int yl = random.nextInt(15);
        g.drawLine(x, y, x + xl, y + yl);
    }

    /**
     *
     * @param num
     * @return Get random characters
     */
    public String getRandomString(int num) {
        return String.valueOf(randString.charAt(num));
    }

    /**
     *Generate random images
     * @param request
     * @param response
     * @param key
     */
    public void getRandcode(HttpServletRequest request, HttpServletResponse response, String key) {

// The BufferedImage class is an Image class with a buffer, and the Image class is a class that describes the image information
        BufferedImage image = new BufferedImage(width, height, BufferedImage.TYPE_INT_BGR);
        Graphics g = image.getGraphics();// 产生Image对象的Graphics对象,改对象可以在图像上进行各种绘制操作
        g.fillRect(0, 0, width, height);
        g.setFont(new Font("Times New Roman", Font.ROMAN_BASELINE, 19));
        g.setColor(getRandColor(110, 133));
// Draw interference lines
        for (int i = 0; i <= lineSize; i++) {
            drowLine(g);
        }
// Draw random characters
        String randomString = "";
        for (int i = 1; i <= stringNum; i++) {
            randomString = drowString(g, randomString, i);
        }
//1：Randomly generated verification code into the cookie
        Cookie cookie = new Cookie(key, randomString);
        response.addCookie(cookie);
//2：Randomly generated verification code into the session
//        String sessionid = request.getSession().getId();
//        request.getSession().setAttribute(sessionid + key, randomString);
//        System.out.println("*************" + randomString);

        g.dispose();
        try {
            ByteArrayOutputStream tmp = new ByteArrayOutputStream();
            ImageIO.write(image, "png", tmp);
            tmp.close();
            Integer contentLength = tmp.size();
            response.setHeader("content-length", contentLength + "");
            response.getOutputStream().write(tmp.toByteArray());// The picture in memory is output to the client in streaming form
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            try {
                response.getOutputStream().flush();
                response.getOutputStream().close();
            } catch (Exception e2) {
                e2.printStackTrace();
            }
        }
    }

}