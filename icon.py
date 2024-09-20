# 本文件由ChatGPT生成，用于生成电池图标
import os
from PIL import Image, ImageDraw, ImageFont

def create_icon(percentage, charging):
    size = (32, 32)
    color = "green" if charging else "red"
    transparent_background = (0, 0, 0, 0)  # 设置透明背景
    
    img = Image.new("RGBA", size, transparent_background)  # 使用透明背景创建图标
    draw = ImageDraw.Draw(img)
    
    # 绘制电量框架，去掉不必要的白色边框
    draw.rectangle([(2, 2), (30, 30)], outline=color, width=2)  # 位置稍微靠近边界以去除白边
    
    # 填充电量
    fill_height = int((percentage / 100) * 26)  # 调整高度以适应新的框架
    draw.rectangle([(4, 30 - fill_height), (28, 30)], fill=color)  # 同样调整填充位置
    
    # 移除闪电符号
    # 根据数字的长度选择合适的字体大小
    try:
        font_size = 20
        font = ImageFont.truetype("arial", font_size)
    except IOError:
        font = ImageFont.load_default()
    
    text = f"{percentage}"
    text_bbox = draw.textbbox((0, 0), text, font=font)
    text_width = text_bbox[2] - text_bbox[0]
    text_height = text_bbox[3] - text_bbox[1]

    # 根据百分比设置字体颜色和位置, 全部改为黄色字体
    text_color = "cyan"
    if percentage > 30:
        text_x = (size[0] - text_width) // 2 + 1
        text_y = 7
    else:
        text_x = (size[0] - text_width) // 2 + 0.4
        text_y = 7

    draw.text((text_x, text_y), text, font=font, fill=text_color)
    
    return img

# 创建目录
os.makedirs("icons/charging", exist_ok=True)
os.makedirs("icons/discharging", exist_ok=True)

# 生成0到100%的图标
for i in range(101):
    img = create_icon(i, charging=True)
    img.save(f"icons/charging/icon_{i}.ico")
    img = create_icon(i, charging=False)
    img.save(f"icons/discharging/icon_{i}.ico")
