import numpy as np
import cv2
import tensorflow as tf
from PIL import Image
from keras.models import load_model

model = load_model('effnet.h5')


def img_pred(path):
    img = Image.open(path)
    import fileinput
    
    opencvImage = cv2.cvtColor(np.array(img), cv2.COLOR_RGB2BGR)
    img = cv2.resize(opencvImage,(150,150))
    img = img.reshape(1,150,150,3)
    p = model.predict(img)
    p = np.argmax(p,axis=1)[0]
    f = open("result.txt","w")

    if p==0:
        p='Glioma Tumor'
    elif p==1:
        f.write('The model predicts that there is no tumor')
    elif p==2:
        p='Meningioma Tumor'
    else:
        p='Pituitary Tumor'

    if p!=1:
        f.write(f'The Model predicts that it is a {p}')
        
path = "C:/Users/Harshal Trivedi/Desktop/AU_Activity/cloud-ser/brain_image.jpg"


img_pred(path)



