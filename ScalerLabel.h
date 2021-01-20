#ifndef SCALERLABEL_H
#define SCALERLABEL_H
#include <QLabel>

enum EmDirection
{
    DIR_TOP = 0,
    DIR_BOTTOM,
    DIR_LEFT,
    DIR_RIGHT,
    DIR_LEFTTOP,
    DIR_LEFTBOTTOM,
    DIR_RIGHTTOP,
    DIR_RIGHTBOTTOM,
    DIR_MIDDLE,
    DIR_NONE
};

class ScalerLabel : public QLabel
{
    Q_OBJECT

public:
    ScalerLabel(QWidget *parent = nullptr);
    ~ScalerLabel();
    void setLocationAndSize(int x, int y, int w, int h, int num);
    void setFocus(int index);
    int getFocus();
    void setSpecifyLayout(int x, int y, int w, int h, int index);
    void setVideoWallMode(bool value);
    void EnableMoveChange(bool move, bool change);
    void rePaint(int Ver, int Hor, int total);
    void rePaint(int num);
    void setOutputResolution(int w, int h);
    void setShowLocationType(int type);
    QRect getFoucsLayout();
    QRect* getFoucsGroupLayout(int group);

    void setFoucsLayout(QRect rect);
protected:
    void paintEvent(QPaintEvent *event);
    virtual void mousePressEvent(QMouseEvent *event);
    virtual void mouseReleaseEvent(QMouseEvent *event);
    virtual void mouseDoubleClickEvent(QMouseEvent *event);
    virtual void mouseMoveEvent(QMouseEvent *event);

private:
    bool m_bPainterPressed;        //
    bool m_bMovedPressed;          //
    bool m_bScalePressed;

    void InitialRectangle();


    int indexFocus = 0;
    int panelX;
    int panelY;
    int panelW;
    int panelH;
    int M;
    int N;
    int Num;
    int realWidth = 1920;
    int realHeight = 1080;
    int globalX;
    int globalY;
    QRect *CusRect;
    QRect *realRect;
    QColor defautColor[4] = {QColor(255,132,16,200),QColor(40,107,204,200),QColor(220,8,0,200),QColor(81,177,29,200)};
    bool useDefaultColor = true;
    int showLocationType = 1; //0: no show, 1: only focus, 2: focus group, 3: all, 5: only foucs group rectangle
    bool fourbyfour = false;
    bool canChange = false;
    bool drawMoveRec = true;
    bool EnableMove = true;
    bool EnableChange = true;
    bool OnlyFocus = false;
    bool canMove = false;
    bool isWall = false;
    int moveType = 0;
    void drawMoveRectangle(QRect Rec);
    void ChangeEachLayout(int Lx, int Ly);
    void TransferFourByFour();
    int isBoundry(int x, int y);
};

#endif // SCALERLABEL_H
